package controllers

import (
	"net/http"
	"fmt"
	"io/ioutil"
	log "github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/context"
	"encoding/json"
	"errors"
	ps "github.com/itsbalamurali/heyasha/core/engine/stt"
	"github.com/unrolled/render"
	"github.com/itsbalamurali/heyasha/controllers/stt"
)

//const MAX_MEMORY = 1 * 1024 * 1024

var renderer = render.New(render.Options{})

func SpeechProcess(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	//w.Header().Set("Connection", "Keep-Alive")
	//w.Header().Set("Transfer-Encoding", "chunked")

	content_type := r.Header.Get("Content-Type")
	fmt.Println(content_type)
	buf, err := ioutil.ReadAll(r.Body)
	if err!=nil {log.Fatal("request",err)}
	fmt.Println(buf)

	/*
	if err := r.ParseMultipartForm(MAX_MEMORY); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusForbidden)
	}

	for key, value := range r.MultipartForm.Value {
		fmt.Fprintf(w, "%s:%s ", key, value)
		log.Printf("%s:%s", key, value)
	}

	for _, fileHeaders := range r.MultipartForm.File {
		for _, fileHeader := range fileHeaders {
			file, _ := fileHeader.Open()
			path := fmt.Sprintf("files/%s", fileHeader.Filename)
			buf, _ := ioutil.ReadAll(file)
			ioutil.WriteFile(path, buf, os.ModePerm)
		}
	}*/
}

type Rsponse struct {
	Response ps.Result `json:"response"`
}

func WriteJsonErrorResponse(w http.ResponseWriter, message string, code int) {
	renderer.JSON(w, code, map[string]string{
		"message": message,
	})
}

func getSphinx(ctx context.Context) (*ps.PsInstance, string, string, error) {
	lang := stt.LangFromContext(ctx)
	sp, ok := ps.FromContext(ctx)
	if !ok {
		log.Errorln("speech recognition engine is not ready")
		return nil, lang, "speech recognition engine is not ready", errors.New("speech recognition engine is not ready")
	}
	ps, err := sp.GetSphinxFromLanguage(lang)
	if err != nil {
		log.Errorln(err)
		return nil, lang, "invalid accept-language", err

	}
	return ps, lang, "", nil
}

func dictationHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	ps, lang, errmsg, err := getSphinx(ctx)
	if err != nil {
		WriteJsonErrorResponse(w, errmsg, 500)
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
		size, err := r.Body.Read(buf)
		if err != nil {
			break
		}
		ps.ProcessRaw(buf[:size], false, false)
	}
	ps.EndUtt()
	res, err := ps.GetHyp(false)
	if err != nil {
		WriteJsonErrorResponse(w, "recognition error", 500)
		return
	}
	bytes, err := json.Marshal(Rsponse{res})
	if err != nil {
		WriteJsonErrorResponse(w, "error", 500)
	}
	renderer.JSON(w, http.StatusOK, bytes)
	return
}