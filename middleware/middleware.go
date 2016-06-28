//Project : Hey Asha!
//Copyright (C) Balamurali Pandranki - All Rights Reserved
//Unauthorized copying of this file, via any medium is strictly prohibited
//Proprietary and confidential
//Written by Balamurali Pandranki <balamurali@live.com>, 3/5/2016 2:16 AM
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"github.com/itsbalamurali/heyasha/core/database"
	log "github.com/Sirupsen/logrus"
	jwt_lib "github.com/dgrijalva/jwt-go"
	"time"
	"github.com/itsbalamurali/heyasha/config"
)


//Api authentication middleware
func respondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}
	c.JSON(code, resp)
	c.AbortWithStatus(code)
}

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		_, err := jwt_lib.ParseFromRequest(c.Request, func(token *jwt_lib.Token) (interface{}, error) {
			b := ([]byte(config.AppSecret))
			return b, nil
		})

		if err != nil {
			c.JSON(401,gin.H{"success":false,"message":"unauthorized"})
			//c.AbortWithError(401, err)
			c.Abort()

		}


		/*
		token := c.Request.Header.Get("X-Auth-Token")
		if len(token) == 0 {
			token = c.Query("auth_token")
		}
		if token == "" {
			respondWithError(401, "API token required", c)
			return
		}
		if token != "mytoken" {
			respondWithError(401, "Invalid API token", c)
			return
		}
		c.Next()*/
	}
}

//Request ID middleware
func RequestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := uuid.NewV4().String()
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Access-Control-Allow-Origin")
		c.Writer.Header().Set("X-Request-Id", uuid)
		c.Writer.Header().Set("X-Asha-Version", "1.0")
		c.Set("x-request-id",uuid)
		c.Next()
	}
}


//Mysql Connection middleware
func MysqlConware() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.Db
		//defer db.Close()
		c.Set("mysql", db.New())
		c.Next()
	}
}

func Ginrus(logger *log.Logger, timeFormat string, utc bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// some evil middlewares modify this values
		path := c.Request.URL.Path
		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		if utc {
			end = end.UTC()
		}

		entry := logger.WithFields(log.Fields{
			"status":     c.Writer.Status(),
			"method":     c.Request.Method,
			"path":       path,
			"ip":         c.ClientIP(),
			"latency":    latency,
			"user-agent": c.Request.UserAgent(),
			"time":       end.Format(timeFormat),
		})

		if len(c.Errors) > 0 {
			// Append error field if this is an erroneous request.
			entry.Error(c.Errors.String())
		} else {
			entry.Info()
		}
	}
}
