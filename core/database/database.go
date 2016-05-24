//Project : Hey Asha!
//Copyright (C) Balamurali Pandranki - All Rights Reserved
//Unauthorized copying of this file, via any medium is strictly prohibited
//Proprietary and confidential
//Written by Balamurali Pandranki <balamurali@live.com>, 25/4/2016 4:27 PM
package database

import (
	"gopkg.in/mgo.v2"
	"github.com/itsbalamurali/heyasha/config"
	"log"
)

var (
	// Session stores mongo session
	Session *mgo.Session
	// Mongo stores the mongodb connection string information
	Mongo *mgo.DialInfo
)

const (
	// MongoDBUrl is the default mongodb url that will be used to connect to the
	// database.
	MongoDBUrl = "mongodb://murali:yourass1994@ds011872.mlab.com:11872/heyasha"
)

// Connect connects to mongodb
func Connect() {
	config := config.LoadConfig()
	uri := config.MongoURI
	if len(uri) == 0 {
		uri = MongoDBUrl
	}
	mongo, err := mgo.ParseURL(uri)
	s, err := mgo.Dial(uri)
	if err != nil {
		log.Printf("Can't connect to mongo, go error %v\n", err.Error())
		//panic(err.Error())
	}
	s.SetSafe(&mgo.Safe{})
	log.Println("Connected to", uri)
	Session = s
	Mongo = mongo
}