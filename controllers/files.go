//Project : Hey Asha!
//Copyright (C) Balamurali Pandranki - All Rights Reserved
//Unauthorized copying of this file, via any medium is strictly prohibited
//Proprietary and confidential
//Written by Balamurali Pandranki <balamurali@live.com>, 17/06/2016 1:47 AM
package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/itsbalamurali/heyasha/models"
	"net/http"
)


//Upload File
func FileUpload(c *gin.Context)  {
	db := c.MustGet("mysql").(*gorm.DB)
	dbfile := &models.File{
		UserID: "",
		OriginalName: "",
		ContentType: "",
		Uuid: "",
	}
	db.Save(&dbfile)


}

//Get File by ID https://api.heyasha.com/v1/files/{uuid}
func FileGetById(c *gin.Context)  {
	db := c.MustGet("mysql").(*gorm.DB)
	file := &models.File{}
	fileid := c.Param("uuid")
	db.Where("uuid = ?",fileid).First(&file)
	redirect := c.Query("redirect")
	if redirect == "true"{
		c.Redirect(http.StatusMovedPermanently,"https://storage.googleapis.com/hey-asha.appspot.com/media/"+file.Uuid)
	} else {
		c.JSON(http.StatusOK,gin.H{"success":true,"link":"https://storage.googleapis.com/hey-asha.appspot.com/media/"+file.Uuid,})
	}
}
