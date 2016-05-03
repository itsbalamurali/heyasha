package controllers

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	ps "github.com/itsbalamurali/heyasha/core/engine/stt"
	"github.com/unrolled/render"
	"golang.org/x/text/language"
	"net/http"
	"runtime"
	"github.com/gin-gonic/gin"
	"github.com/itsbalamurali/heyasha/config"
)

type langkey int

type Rsponse struct {
	Response ps.Result `json:"response"`
}

var (
	englishBase, _         = language.AmericanEnglish.Base()
	lkey           langkey = 0
)

var renderer = render.New(render.Options{})

func SpeechProcess(c *gin.Context) {
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
	ps, lang, errmsg, err := getSphinx(lang)
	if err != nil {
		//WriteJsonErrorResponse(w, errmsg, 500)
		c.JSON(http.StatusInternalServerError,errmsg)
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
		c.JSON(http.StatusInternalServerError,errmsg)
		//WriteJsonErrorResponse(w, "recognition error", 500)
		return
	}
	bytes, err := json.Marshal(Rsponse{res})
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":"error"})

		//WriteJsonErrorResponse(w, "error", 500)
	}
	//renderer.JSON(w, http.StatusOK, bytes)
	c.JSON(http.StatusOK,bytes)
	return

}

func WriteJsonErrorResponse(w http.ResponseWriter, message string, code int) {
	renderer.JSON(w, code, map[string]string{
		"message": message,
	})
}

func getSphinx(lang string) (*ps.PsInstance, string, string, error) {
	cpus := runtime.NumCPU()
	psAll := ps.NewSphinx(config.LoadConfig().Pocketsphinx, cpus)
	sp := psAll
	/*
		if sp != n {
			log.Errorln("speech recognition engine is not ready")
			return nil, lang, "speech recognition engine is not ready", errors.New("speech recognition engine is not ready")
		}*/
	ps, err := sp.GetSphinxFromLanguage(lang)
	if err != nil {
		//log.Errorln(err)
		return nil, lang, "invalid accept-language", err
	}
	return ps, lang, "", nil
}
