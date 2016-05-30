//Project : Hey Asha!
//Copyright (C) Balamurali Pandranki - All Rights Reserved
//Unauthorized copying of this file, via any medium is strictly prohibited
//Proprietary and confidential
//Written by Balamurali Pandranki <balamurali@live.com>, 9/5/2016 10:39 AM
package models

import "github.com/jinzhu/gorm"

type ConversationLog struct {
	gorm.Model
	Input string
	Response string
	UserID int
	ConvoID string
}