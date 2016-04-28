//Project : Hey Asha!
//Copyright (C) Balamurali Pandranki - All Rights Reserved
//Unauthorized copying of this file, via any medium is strictly prohibited
//Proprietary and confidential
//Written by Balamurali Pandranki <balamurali@live.com>, 25/4/2016 4:27 PM
package main

import (
	"github.com/itsbalamurali/heyasha/controllers"
	"github.com/itsbalamurali/heyasha/controllers/platforms"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
	"runtime"
	"github.com/itsbalamurali/heyasha/shared/datatypes"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
)

func main() {
	// maximize CPU usage for maximum performance
	runtime.GOMAXPROCS(runtime.NumCPU())

	//Init Logger
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stderr)

	env := os.Getenv("ENV")
	if *env == "production" {
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
	router := httprouter.New()

	//Hello!!
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		//w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(datatypes.ApiDefaultResponse{Version:"1.0",Message:"Hello, I'm listening!"})
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

	//Communication Platforms
	router.POST("/chat/slack", platforms.SlackBot)         //SlackBot
	router.POST("/chat/kik", platforms.KikBot)             //Kik Bot
	router.POST("/chat/telegram", platforms.TelegramBot)   //Telegram Bot
	router.POST("/chat/skype", platforms.SkypeBot)         //Skype Bot
	router.POST("/chat/messenger", platforms.MessengerBot) //Messenger Bot
	router.GET("/chat/messenger", platforms.MessengerBot)  //Facebook Callback Verification
	router.POST("/chat/sms", platforms.SmsBot)             //Sms Bot
	router.POST("/chat/email", platforms.EmailBot)         //Email Bot

	// Start server
	log.Infoln("Hi, I am running on port: " + port + " !!")
	log.Fatalln(http.ListenAndServe(":"+port, router))
}
