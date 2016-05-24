//Project : Hey Asha!
//Copyright (C) Balamurali Pandranki - All Rights Reserved
//Unauthorized copying of this file, via any medium is strictly prohibited
//Proprietary and confidential
//Written by Balamurali Pandranki <balamurali@live.com>, 3/5/2016 2:16 AM
package middleware

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"time"
	"github.com/sebest/logrusly"
)

type ErrorType uint64

const (
	ErrorTypeBind    ErrorType = 1 << 63 // used when c.Bind() fails
	ErrorTypeRender  ErrorType = 1 << 62 // used when c.Render() fails
	ErrorTypePrivate ErrorType = 1 << 0
	ErrorTypePublic  ErrorType = 1 << 1

	ErrorTypeAny ErrorType = 1<<64 - 1
	ErrorTypeNu            = 2
)


//Api authentication middleware
func respondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}
	c.JSON(code, resp)
	c.AbortWithStatus(code)
}

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//token := c.Request.Header.Get("X-Auth-Token")
		token := c.Query("auth_token")

		if token == "" {
			respondWithError(401, "API token required", c)
			return
		}

		if token != "mytoken" {
			respondWithError(401, "Invalid API token", c)
			return
		}

		c.Next()
	}
}

//Request ID middleware
func RequestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Request-Id", uuid.NewV4().String())
		c.Next()
	}
}

func GinLogger(notlogged ...string) gin.HandlerFunc {
	log.SetFormatter(&log.JSONFormatter{})
	hook := logrusly.NewLogglyHook("09af9fc7-1db3-4c39-a452-f923467e3af1", "heyasha.loggly.com", log.InfoLevel)
	log.AddHook(hook)
	var skip map[string]struct{}

	if length := len(notlogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, path := range notlogged {
			skip[path] = struct{}{}
		}
	}

	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path

		// Process request
		c.Next()
		// Log only when path is not being skipped
		if _, ok := skip[path]; !ok {
			// Stop timer
			end := time.Now()
			latency := end.Sub(start)

			clientIP := c.ClientIP()
			method := c.Request.Method
			statusCode := c.Writer.Status()
			comment := c.Errors.ByType(gin.ErrorTypePrivate).String()

			log.WithFields(log.Fields{
				"time":      end.Format("2006/01/02 - 15:04:05"),
				"status":    statusCode,
				"latency":   latency,
				"ip": clientIP,
				"method":    method,
				"path":      path,
			}).Debug(comment)


			//Flush loggly hook
			// Flush is automatic for panic/fatal
			// Just make sure to Flush() before exiting or you may loose up to 5 seconds
			// worth of messages.
			hook.Flush()
		}

	}
}
