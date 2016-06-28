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
		port = "80"
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

	router.LoadHTMLGlob("../assets/html/web/*")

	//Hello!!
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{"title": "Asha - Your Artificial Intelligent Friend ;)"})
	})

	router.GET("/chat", func(c *gin.Context) {
		c.HTML(http.StatusOK, "chat.tmpl", gin.H{"title": "Asha - Your Artificial Intelligent Friend ;)"})
	})

	router.GET("/privacy", func(c *gin.Context) {
		c.HTML(http.StatusOK, "privacy.tmpl", gin.H{"title": "Asha - Your Artificial Intelligent Friend ;)"})
	})

	router.GET("/terms-of-service", func(c *gin.Context) {
		c.HTML(http.StatusOK, "terms.tmpl", gin.H{"title": "Asha - Your Artificial Intelligent Friend ;)"})
	})

	router.GET("/app", func(c *gin.Context) {
		c.HTML(http.StatusOK, "app.tmpl", gin.H{"title": "Asha - Your Artificial Intelligent Friend ;)"})
	})

	router.GET("/purpose", func(c *gin.Context) {
		c.HTML(http.StatusOK, "purpose.tmpl", gin.H{"title": "Asha - Your Artificial Intelligent Friend ;)"})
	})

	router.GET("/faq", func(c *gin.Context) {
		c.HTML(http.StatusOK, "faq.tmpl", gin.H{"title": "Asha - Your Artificial Intelligent Friend ;)"})
	})

	router.GET("/blog", func(c *gin.Context) {
		c.HTML(http.StatusOK, "blog.tmpl", gin.H{"title": "Asha - Your Artificial Intelligent Friend ;)"})
	})

	router.GET("/blog/:title", func(c *gin.Context) {
		c.HTML(http.StatusOK, "blog_post.tmpl", gin.H{"title": "Asha - Your Artificial Intelligent Friend ;)"})
	})

	//404 Handler
	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "404.tmpl", gin.H{"title": "Asha - Your Artificial Intelligent Friend ;)"})
	})

	//Start server
	log.Println("Web server is running!!")
	manners.ListenAndServe(":"+port, router) //Graceful restarts

}
