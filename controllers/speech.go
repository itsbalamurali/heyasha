package controllers

import (
	"github.com/gin-gonic/gin"
)

func SpeechProcess(c *gin.Context) {

	buf := make([]byte, 1024)
	// Detect Language from accept-language
	lang := c.Request.Header.Get("Accept-Language")

	for {
		size, err := c.Request.Body.Read(buf)
		if err != nil {
			break
		}
		ps.ProcessRaw(buf[:size], false, false)
	}
}
