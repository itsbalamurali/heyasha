//Project : Hey Asha!
//Copyright (C) Balamurali Pandranki - All Rights Reserved
//Unauthorized copying of this file, via any medium is strictly prohibited
//Proprietary and confidential
//Written by Balamurali Pandranki <balamurali@live.com>, 6/5/2016 7:02 PM
package models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Pid string `json:"pid" bson:"pid"`
	Username  string `json:"username" bson:"username"`
	ProfilePicURL string `json:"profile_pic" bson:"profile_pic"`
	Password  string `json:"password" bson:"password"`
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Gender    string `json:"gender" bson:"gender"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	AuthData  Authdata `json:"auth_data"`
	Platforms []Platform `json:"platforms"`
	CreatedOn int64         `json:"created_on" bson:"created_on"`
	UpdatedOn int64         `json:"updated_on" bson:"updated_on"`

}

type Platform struct {
	PlatformID string `json:"profile_id"`
	Name string `json:"name"`
}


type Authdata struct {
	Facebook  FacebookAuthData `json:"facebook"`
	Twitter   TwitterAuthData `json:"twitter"`
	Anonymous AnonymousAuthData `json:"anonymous"`
}

type AnonymousAuthData struct {
	Id string `json:"id"`
}

type FacebookAuthData struct {
	Id          string `json:"id"`
	AccessToken string `json:"access_token"`
	ExpiryDate  string `json:"expiry_date"`
}

type TwitterAuthData struct {
	Id              string `json:"id"`
	ScreenName      string `json:"screen_name"`
	ConsumerKey     string `json:"consumer_key"`
	ConsumerSecret  string `json:"consumer_secret"`
	AuthToken       string `json:"auth_token"`
	AuthTokenSecret string `json:"auth_token_secret"`
}