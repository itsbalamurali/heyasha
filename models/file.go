//Project : Hey Asha!
//Copyright (C) Balamurali Pandranki - All Rights Reserved
//Unauthorized copying of this file, via any medium is strictly prohibited
//Proprietary and confidential
//Written by Balamurali Pandranki <balamurali@live.com>, 6/5/2016 7:02 PM
package models
import "github.com/jinzhu/gorm"

type File struct {
	gorm.Model
	UserID uint64
	OriginalName string
	ContentType string
	Uuid string
}