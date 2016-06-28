//Project : Hey Asha!
//Copyright (C) Balamurali Pandranki - All Rights Reserved
//Unauthorized copying of this file, via any medium is strictly prohibited
//Proprietary and confidential
//Written by Balamurali Pandranki <balamurali@live.com>, 17/06/2016 1:47 AM
package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/itsbalamurali/heyasha/models"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	storage "google.golang.org/api/storage/v1"
	"log"
	"net/http"
	"os"
	"github.com/satori/go.uuid"
	"io"
	"mime"
)

const (
	scope = storage.DevstorageFullControlScope
)


//Upload File
func FileUpload(c *gin.Context) {
	fileType := c.Request.Header.Get("Content-Type")
	fileName := uuid.NewV4().String()
	//fileObject := c.Request.Body
	//buf, _ := ioutil.ReadAll(fileObject)
	//fileExt := http.DetectContentType(buf)
	ftype,err := mime.ExtensionsByType(fileType)
	file, err := os.Create(os.TempDir()+fileName+ ftype[0])
	if err != nil {
		log.Fatalf("Unable to create the file: %v", err)
	}
	//Copy content to file
	_, err = io.Copy(file, c.Request.Body)
	if err != nil {
		log.Println("File Copy Error")
	}

	client, err := google.DefaultClient(context.Background(), scope)
	if err != nil {
		log.Fatalf("Unable to get default client: %v", err)
	}

	service, err := storage.New(client)
	if err != nil {
		log.Fatalf("Unable to create storage service: %v", err)
	}
	object := &storage.Object{Name: "media/"+fileName}

	fileUpload, err := os.Open(os.TempDir()+fileName+ ftype[0])
	if err != nil {
		log.Fatalf("Error opening %q: %v", fileType, err)
	}
	if res, err := service.Objects.Insert("hey-asha.appspot.com", object).Media(fileUpload).PredefinedAcl("publicRead").Do(); err == nil {
		//fmt.Printf("Created object %v at location %v\n\n", res.Name, res.SelfLink)
		db := c.MustGet("mysql").(*gorm.DB)
		dbfile := &models.File{
			UserID:       1,
			OriginalName: fileName,
			ContentType:  fileType,
			Uuid:         fileName,
		}

		db.Save(&dbfile)
		c.JSON(http.StatusOK, gin.H{"success": true, "link": "https://storage.googleapis.com/hey-asha.appspot.com/"+res.Name})
	} else {
		log.Fatalf("Objects.Insert failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "please try again!"})
	}

}

//Get File by ID https://api.heyasha.com/v1/files/{uuid}
func FileGetById(c *gin.Context) {
	db := c.MustGet("mysql").(*gorm.DB)
	file := &models.File{}
	fileid := c.Param("uuid")
	db.Where("uuid = ?", fileid).First(&file)
	redirect := c.Query("redirect")
	if redirect == "true" {
		c.Redirect(http.StatusMovedPermanently, "https://storage.googleapis.com/hey-asha.appspot.com/media/"+file.Uuid)
	} else {
		c.JSON(http.StatusOK, gin.H{"success": true, "link": "https://storage.googleapis.com/hey-asha.appspot.com/media/" + file.Uuid})
	}
}