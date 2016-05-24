//Project : Hey Asha!
//Copyright (C) Balamurali Pandranki - All Rights Reserved
//Unauthorized copying of this file, via any medium is strictly prohibited
//Proprietary and confidential
//Written by Balamurali Pandranki <balamurali@live.com>, 6/5/2016 7:02 PM
package models

type User struct {
	Username  string
	Password  string
	FirstName string
	LastName  string
	Gender    string
	Phone     string
	Email     string
	AuthData  Authdata
}

type Authdata struct {
	Facebook  FacebookAuthData
	Twitter   TwitterAuthData
	Anonymous AnonymousAuthData
}

type AnonymousAuthData struct {
	Id string
}

type FacebookAuthData struct {
	Id          string
	AccessToken string
	ExpiryDate  string
}

type TwitterAuthData struct {
	Id              string
	ScreenName      string
	ConsumerKey     string
	ConsumerSecret  string
	AuthToken       string
	AuthTokenSecret string
}
