//Admin Panel Server
/*
* @Author: Balamurali Pandranki
* @Date:   2016-06-29 03:27:52
* @Last Modified by:   Balamurali Pandranki
* @Last Modified time: 2016-06-29 03:31:02
 */

package main

import (
	"net/http"
	"os"
	"runtime"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/braintree/manners"
	"github.com/gin-gonic/gin"
	"github.com/itsbalamurali/heyasha/middleware"
)

func main() {

	// maximize CPU usage for maximum performance
	runtime.GOMAXPROCS(runtime.NumCPU())

	app_env := os.Getenv("ENV")
	if app_env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	//Port to Bind server to
	port := os.Getenv("PORT")
	if port == "" {
		port = "1234"
	}

	//New Router
	router := gin.New()
	// Global middleware
	router.Use(middleware.Ginrus(log.StandardLogger(), time.RFC3339, true))
	router.Use(gin.Recovery())
	router.Use(middleware.RequestIdMiddleware())

	//Static files
	router.Static("/assets", "./assets")

	//Favicon
	router.StaticFile("/favicon.ico", "../assets/img/favicon.png")

	router.LoadHTMLGlob("../assets/html/admin/*")

	//Routes!!
	router.GET("/")
	router.GET("/login")
	router.GET("/users")
	router.POST("/users")
	router.GET("/users/:id")
	router.POST("/users/:id")
	router.DELETE("/users/:id")
	router.GET("/conversations")
	router.GET("/conversations/:id")
	router.GET("/aiml")
	router.GET("/intents")
	router.GET("/intents/:id")
	router.GET("/entities")
	router.GET("/entities/:id")
	router.GET("/plugins")
	router.GET("/plugins/:id")
	router.GET("/blog")
	router.POST("/blog")
	router.GET("/blog/edit/:id")
	router.POST("/blog/edit/:id")

	//404 Handler
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"success": "false", "message": "Asha - 404 Not Found!!"})
	})

	//Start server
	log.Println("Admin server is running!!")
	manners.ListenAndServe(":"+port, router) //Graceful restarts
}
