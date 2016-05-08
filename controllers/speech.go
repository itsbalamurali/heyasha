package controllers

import (
	"github.com/gin-gonic/gin"
)

func SpeechProcess (c *gin.Context) {
	/*fn := func(c *gin.Context) {
		// Detect Language from accept-language
		lang := ""
		if tags, _, err := language.ParseAcceptLanguage(c.Request.Header.Get("Accept-Language")); err == nil {
			if len(tags) > 0 {
				t := tags[0]
				base, _, _ := t.Raw() // base, sscript, region
				if base == englishBase {
					lang = "en-US"
				} else if t == language.BritishEnglish {
					lang = "en-GB"
				}
			}
		}
		log.Infoln("Requested lang: " + lang)
		ps, lang, errmsg, err := getSphinx(lang, ctx)
		if err != nil {
			//WriteJsonErrorResponse(w, errmsg, 500)
			c.JSON(http.StatusInternalServerError, errmsg)
			return
		}

		log.WithFields(log.Fields{
			"lang": lang,
		}).Info()

		log.Infoln(lang)
		ps.Lock()
		defer ps.Unlock()

		buf := make([]byte, 1024)
		ps.StartUtt()
		for {
			size, err := c.Request.Body.Read(buf)
			if err != nil {
				break
			}
			ps.ProcessRaw(buf[:size], false, false)
		}
		ps.EndUtt()
		res, err := ps.GetHyp(false)
		if err != nil {
			c.JSON(http.StatusInternalServerError, errmsg)
			//WriteJsonErrorResponse(w, "recognition error", 500)
			return
		}
		bytes, err := json.Marshal(Rsponse{res})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error"})

			//WriteJsonErrorResponse(w, "error", 500)
		}
		//renderer.JSON(w, http.StatusOK, bytes)
		c.JSON(http.StatusOK, bytes)
		return
	}
	return gin.HandlerFunc(fn)*/
}