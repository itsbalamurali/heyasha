//Project : Hey Asha!
//Copyright (C) Balamurali Pandranki - All Rights Reserved
//Unauthorized copying of this file, via any medium is strictly prohibited
//Proprietary and confidential
//Written by Balamurali Pandranki <balamurali@live.com>, 3/5/2016 2:16 AM
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
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