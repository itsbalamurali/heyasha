package main

import (
	"net/http"
	"os"
	"log"
	"github.com/itsbalamurali/bot/controllers"
	"github.com/itsbalamurali/bot/controllers/platforms"
	"github.com/julienschmidt/httprouter"
	"runtime"
	"fmt"
)

const MAX_UPLOAD_MEMORY = 1 * 1024 * 1024

func main() {

	// maximize CPU usage for maximum performance
	runtime.GOMAXPROCS(runtime.NumCPU())

	//Port to Bind server to
	port := os.Getenv("PORT")
	if port == ""{
		port = "80"
	}

	//New Router
	router := httprouter.New()


	//Hello!!
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Welcome!\n")
	})
	//Core REST API routes
	router.POST("/speech", controllers.AudioUpload) //speech recognition
	router.GET("/message", controllers.Chat)  //chat with bot
	router.GET("/extract", controllers.IntentExtract) //Extract Intent from Text

	//User REST API routes
	router.POST("/users/", controllers.CreateUser)
	router.POST("/users/login", controllers.LoginUser)
	router.GET("/users/{UserId}",controllers.GetUserDetails)
	router.DELETE("/users/{UserId}", controllers.DeleteUser)

	//Communication Platforms
	router.POST("/chat/slack",platforms.SlackBot) 	    //SlackBot
	router.POST("/chat/kik", platforms.KikBot)             //Kik Bot
	router.POST("/chat/telegram", platforms.TelegramBot)   //Telegram Bot
	router.POST("/chat/skype", platforms.SkypeBot)         //Skype Bot
	router.POST("/chat/messenger", platforms.MessengerBot) //Messenger Bot
	router.GET("/chat/messenger", platforms.VerifyMessengerBot) //Facebook Callback Verification
	router.POST("/chat/sms", platforms.SmsBot)             //Sms Bot
	router.POST("/chat/email", platforms.EmailBot) //Email Bot

	// Start server
	fmt.Println("Hi, I am running on port: "+ port +" !!")
	log.Fatal(http.ListenAndServe(":"+ port, router))

}
