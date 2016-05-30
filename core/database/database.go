//Project : Hey Asha!
//Copyright (C) Balamurali Pandranki - All Rights Reserved
//Unauthorized copying of this file, via any medium is strictly prohibited
//Proprietary and confidential
//Written by Balamurali Pandranki <balamurali@live.com>, 25/4/2016 4:27 PM
package database

import (
	"gopkg.in/mgo.v2"
	"os"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/itsbalamurali/heyasha/models"
)

var (
	// Session stores mongo session
	Session *mgo.Session
	// Mongo stores the mongodb connection string information
	Mongo *mgo.DialInfo
	//Mysql Connection
	Db *gorm.DB
)

const (
	// MongoDBUrl is the default mongodb url that will be used to connect to the
	// database.
	MongoDBUrl = "mongodb://murali:yourass1994@ds011872.mlab.com:11872/heyasha"
)

// Connect connects to mongodb
func Connect() {
	uri := os.Getenv("MONGODB_URL")

	if len(uri) == 0 {
		uri = MongoDBUrl
	}

	mongo, err := mgo.ParseURL(uri)
	s, err := mgo.Dial(uri)
	if err != nil {
		log.Printf("Can't connect to mongo, go error %v\n", err)
		panic(err.Error())
	}
	s.SetSafe(&mgo.Safe{})
	log.Println("Connected to", uri)
	Session = s
	Mongo = mongo
}

func ConnectMysql()  {
	db, err := gorm.Open("mysql", "admin_twa:Aydim4VdBK@tcp(128.199.81.183:3306)/admin_twa?charset=utf8&parseTime=True")
	if err != nil {
		log.Printf("Can't connect to MySQL, go error %v\n", err)
		panic(err.Error())
	}
	db.AutoMigrate(&models.ConversationLog{})
	log.Println("Connected to MySQL Server")
	Db = db
}