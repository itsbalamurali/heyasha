//Project : Hey Asha!
//Copyright (C) Balamurali Pandranki - All Rights Reserved
//Unauthorized copying of this file, via any medium is strictly prohibited
//Proprietary and confidential
//Written by Balamurali Pandranki <balamurali@live.com>, 25/4/2016 4:27 PM
package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/itsbalamurali/heyasha/controllers"
	"github.com/itsbalamurali/heyasha/controllers/platforms"
	"net/http"
	"os"
	"runtime"
	"github.com/itsbalamurali/heyasha/middleware"
)

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


	//Hello!!
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"version": "1.0", "message": "Hello, I'm listening!"})
	})

	//Core REST API routes
	router.POST("/speech", controllers.SpeechProcess)    //speech recognition
	router.GET("/message", controllers.Chat)           //chat with bot
	router.GET("/extract", controllers.IntentExtract)  //Extract Intent from Text
	router.GET("/suggest", controllers.SuggestQueries) //Autocomplete user queries

	//User REST API routes
	router.POST("/users/", controllers.CreateUser)
	router.POST("/users/login", controllers.LoginUser)
	router.PUT("/users/{UserId}", controllers.UpdateUserDetails)
	router.GET("/users/{UserId}", controllers.GetUserDetails)
	router.DELETE("/users/{UserId}", controllers.DeleteUser)

	//Sync Adapters
	//router.POST("/sync/contacts")
	//router.POST("/sync/calender")
	//router.POST("/sync/notes")

	//Communication Platforms
	router.POST("/chat/slack", platforms.SlackBot)         //SlackBot
	router.POST("/chat/kik", platforms.KikBot)             //Kik Bot
	router.POST("/chat/telegram", platforms.TelegramBot)   //Telegram Bot
	router.POST("/chat/skype", platforms.SkypeBot)         //Skype Bot
	router.POST("/chat/messenger", platforms.MessengerBot) //Messenger Bot
	router.GET("/chat/messenger", platforms.MessengerBot)  //Facebook Callback Verification
	router.POST("/chat/sms", platforms.SmsBot)             //Sms Bot
	router.POST("/chat/email", platforms.EmailBot)         //Email Bot


	//Method not allowed
	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"message": "method not allowed"} )
	})

	//404 Handler
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "method not found"})
	})

	// Start server
	log.Infoln("Hi, I am running on port: " + port + " !!")
	log.Infoln(router.Run(":" + port))
}
