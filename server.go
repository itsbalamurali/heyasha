//Project : Hey Asha!
//Copyright (C) Balamurali Pandranki - All Rights Reserved
//Unauthorized copying of this file, via any medium is strictly prohibited
//Proprietary and confidential
//Written by Balamurali Pandranki <balamurali@live.com>, 25/4/2016 4:27 PM
package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/itsbalamurali/heyasha/controllers"
	"github.com/itsbalamurali/heyasha/controllers/platforms"
	"github.com/itsbalamurali/heyasha/middleware"
	"os"
	"runtime"
	"net/http"
	"github.com/sebest/logrusly"
)

var logglyToken string = "09af9fc7-1db3-4c39-a452-f923467e3af1"

//var engine *xorm.Engine
func init() {

}

func main() {
	// maximize CPU usage for maximum performance
	runtime.GOMAXPROCS(runtime.NumCPU())

	//Init Logger
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stderr)

	app_env := os.Getenv("ENV")
	if app_env == "production" {
		log.SetLevel(log.WarnLevel)
	} else {
		log.SetLevel(log.DebugLevel)
	}
	hook := logrusly.NewLogglyHook(logglyToken, "www.hostname.com", log.WarnLevel, "tag1", "tag2")
	log.AddHook(hook)
	log.Infoln("Starting server...")

	//Database error variable and engine
	//var err error
	//engine, err = xorm.NewEngine("mysql", "root:123@/test?charset=utf8")
	//logger := xorm.NewSimpleLogger(log.Logger{})
	//logger.ShowSQL(true)
	//engine.SetLogger(logger)

	//Port to Bind server to
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	//New Router
	router := gin.Default()

	// Global middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.RequestIdMiddleware())
	router.Use(middleware.TokenAuthMiddleware())

	//Hello!!
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"version": "1.0", "message": "Hello, I'm listening!"})
	})

	//Core REST API routes
	router.POST("/v1/speech", controllers.SpeechProcess)  //speech recognition
	router.GET("/v1/chat", controllers.Chat)              //chat with bot
	router.GET("/v1/extract", controllers.IntentExtract)  //Extract Intent from Text
	router.GET("/v1/suggest", controllers.SuggestQueries) //Autocomplete user queries

	//User REST API routes
	router.POST("/v1/users/", controllers.CreateUser)
	router.POST("/v1/users/login", controllers.LoginUser)
	router.POST("/v1/users/logout", controllers.LoginUser)
	router.GET("/v1/users/{UserId}", controllers.GetUserDetails)
	router.GET("/v1/users/me", controllers.GetUserDetails)
	router.PUT("/v1/users/{UserId}", controllers.UpdateUserDetails)
	router.DELETE("/v1/users/{UserId}", controllers.DeleteUser)
	router.DELETE("/v1/users/reset_password", controllers.DeleteUser)

	//TODO Sessions & Files

	//Sync Adapters
	router.POST("/v1/sync/contacts")
	router.POST("/sync/calender")
	router.POST("/v1/sync/notes")

	//Communication Platforms
	router.POST("/v1/chat/slack", platforms.SlackBot)         //SlackBot
	router.POST("/v1/chat/kik", platforms.KikBot)             //Kik Bot
	router.POST("/v1/chat/telegram", platforms.TelegramBot)   //Telegram Bot
	router.POST("/v1/chat/skype", platforms.SkypeBot)         //Skype Bot
	router.POST("/v1/chat/messenger", platforms.MessengerBot) //Messenger Bot
	router.GET("/v1/chat/messenger", platforms.MessengerBot)  //Facebook Callback Verification
	router.POST("/v1/chat/sms", platforms.SmsBot)             //Sms Bot
	router.POST("/v1/chat/email", platforms.EmailBot)         //Email Bot

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
	log.Infoln(router.Run(":" + port))

	//Flush loggly hook
	// Flush is automatic for panic/fatal
	// Just make sure to Flush() before exiting or you may loose up to 5 seconds
	// worth of messages.
	hook.Flush()
}
