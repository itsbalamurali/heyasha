//Project : Hey Asha!
//Copyright (C) Balamurali Pandranki - All Rights Reserved
//Unauthorized copying of this file, via any medium is strictly prohibited
//Proprietary and confidential
//Written by Balamurali Pandranki <balamurali@live.com>, 6/5/2016 7:02 PM
package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Pid string `json:"pid" bson:"pid"`
	Username  string `json:"username" binding:"required"`
	ProfilePicURL string `json:"profile_pic"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	Phone     string `json:"phone"`
	Email     string `json:"email" binding:"required"`
	AuthData  UsersAuthdata `json:"auth_data"`
	Platforms []UsersPlatform `json:"platforms"`
}

type UsersPlatform struct {
	PlatformID string `json:"profile_id"`
	Name string `json:"name"`
}


type UsersAuthdata struct {
	Facebook  UsersFacebookAuthData `json:"facebook"`
	Twitter   UsersTwitterAuthData `json:"twitter"`
	Anonymous UsersAnonymousAuthData `json:"anonymous"`
}

type UsersAnonymousAuthData struct {
	ID string `json:"id"`
}

type UsersFacebookAuthData struct {
	ID          string `json:"id"`
	AccessToken string `json:"access_token"`
	ExpiryDate  string `json:"expiry_date"`
}

type UsersTwitterAuthData struct {
	Id              string `json:"id"`
	ScreenName      string `json:"screen_name"`
	ConsumerKey     string `json:"consumer_key"`
	ConsumerSecret  string `json:"consumer_secret"`
	AuthToken       string `json:"auth_token"`
	AuthTokenSecret string `json:"auth_token_secret"`
}

type TokenResponse struct {
	Token string `json:"token"`
}