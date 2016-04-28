package controllers

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"log"
	"github.com/julienschmidt/httprouter"
)

const MAX_MEMORY = 1 * 1024 * 1024

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
